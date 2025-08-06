import { AxiosError } from 'axios'
import * as cheerio from 'cheerio'
import * as CryptoJS from 'crypto-js'
import * as Tesseract from 'tesseract.js'
import * as fs from 'fs'
import * as path from 'path'
import request, { getCookie } from '../../shared/request'
import { IQuery } from '../../server'

// 复用验证码识别函数
async function recognizeCaptcha(imageBuffer: Buffer): Promise<string> {
  console.log('[验证码识别] 开始识别验证码')
  
  try {
    const worker = await Tesseract.createWorker('eng')
    await worker.setParameters({
      tessedit_char_whitelist: 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    })

    // 保存验证码图片用于调试
    const debugDir = path.join(__dirname, '..', 'debug')
    if (!fs.existsSync(debugDir)) {
      fs.mkdirSync(debugDir, { recursive: true })
    }
    const imagePath = path.join(debugDir, `captcha-${Date.now()}.png`)
    fs.writeFileSync(imagePath, imageBuffer)

    const { data: { text } } = await worker.recognize(imageBuffer)
    await worker.terminate()
    
    const cleanedText = text.replace(/\s+|\W+/g, '').slice(0, 4) || 'ABCD'
    
    return cleanedText
  } catch (error) {
    console.error('[验证码识别] OCR识别错误:', error)
    return 'ABCD' // 默认值防止空结果
  }
}

const url1 = 'https://webvpn.hfut.edu.cn/http/77726476706e69737468656265737421f3f652d22f367d44300d8db9d6562d/cas/login?service=https%3A%2F%2Fwebvpn.hfut.edu.cn%2Flogin%3Fcas_login%3Dtrue'
const url2 = 'https://webvpn.hfut.edu.cn/wengine-vpn/input'
const url3 = 'https://webvpn.hfut.edu.cn/http/77726476706e69737468656265737421f3f652d22f367d44300d8db9d6562d/cas/checkInitVercode?vpn-12-o1-cas.hfut.edu.cn='
const url4 = 'https://webvpn.hfut.edu.cn/wengine-vpn/cookie?method=get&host=cas.hfut.edu.cn&scheme=http&path=/cas/login'
const vercodeUrl = 'https://webvpn.hfut.edu.cn/http/77726476706e69737468656265737421f3f652d22f367d44300d8db9d6562d/cas/vercode?vpn-12-o1-cas.hfut.edu.cn=&'
const loginUrl = 'https://webvpn.hfut.edu.cn/http/77726476706e69737468656265737421f3f652d22f367d44300d8db9d6562d/cas/login?service=https%3A%2F%2Fcas.hfut.edu.cn%2Fcas%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3DBsHfutEduPortal%26redirect_uri%3Dhttps%253A%252F%252Fone.hfut.edu.cn%252Fhome%252Findex%26response_type%3Dcode%26client_name%3DCasOAuthClient'

function encryptionPwd(pwd: string, salt: string) {
  const key = CryptoJS.enc.Utf8.parse(salt)
  const password = CryptoJS.enc.Utf8.parse(pwd)
  const encrypted = CryptoJS.AES.encrypt(password, key, { 
    mode: CryptoJS.mode.ECB, 
    padding: CryptoJS.pad.Pkcs7 
  })
  return encrypted.toString()
}

export default async function(query: IQuery) {
  const { req } = query
  const username = req.query.username as string
  const password = req.query.password as string

  if (!username || !password) {
    return {
      code: 400,
      msg: '用户名或密码不能为空',
    }
  }

  const maxAttempts = 3
  let attempts = 0
  let success = false
  let result: any = null
  let cookie1 = ''
  let payload: any = {}
  let saltKey = ''
  let encryptedPwd = ''

  while (attempts < maxAttempts && !success) {
    try {
      // 第一步: 获取初始页面和Cookie
      const res1 = await request(url1, {}, query)
      cookie1 = getCookie(res1.cookie as string[]) || ''
      payload = { cookie: cookie1 }

      if (!cookie1) {
        return { code: 400, msg: '无法获取初始Cookie' }
      }

      // 第二步: 初始化VPN
      await request(url2, { params: { _: Date.now() }, maxRedirects: 5 }, payload)
      await request(url3, {}, payload)
      
      // 第三步: 获取加密盐值
      const res4 = await request(url4, {}, payload)
      saltKey = (res4.body as string).split('; ')[1]?.split('=').pop() || ''
      encryptedPwd = encryptionPwd(password, saltKey)
      
      // 第四步: 获取验证码并识别
      console.log('获取验证码...')
      const vercodeRes = await request(vercodeUrl, {
        responseType: 'arraybuffer',
        params: { _: Date.now() }
      }, payload)
      
      const captchaText = await recognizeCaptcha(vercodeRes.body)
      console.log(`识别验证码: ${captchaText}`)
      
      // 第五步: 验证用户身份
      const authFlagUrl = `https://webvpn.hfut.edu.cn/http/77726476706e69737468656265737421f3f652d22f367d44300d8db9d6562d/cas/policy/checkUserIdenty?vpn-12-o1-cas.hfut.edu.cn=&username=${username}&password=${encryptedPwd}&capcha=${captchaText}&_=${Date.now()}`
      await request(authFlagUrl, {}, payload)
      
      // 第六步: 提交登录请求
      console.log('提交登录请求...')
      const loginParams = new URLSearchParams({
        username,
        capcha: captchaText,
        execution: 'e1s1',
        _eventId: 'submit',
        password: encryptedPwd,
        geolocation: '',
      }).toString()
      
      let redirectRes: any
      try {
        redirectRes = await request(loginUrl, {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
            'Cookie': payload.cookie
          },
          data: loginParams,
          maxRedirects: 5
        }, payload)
      } catch (err) {
        redirectRes = { 
          body: (err as AxiosError).response?.data || '' 
        }
      }
      
      // 第七步: 检查登录结果
      const $ = cheerio.load(redirectRes.body || '')
      
      if ($('.alert-danger').children('span').text().trim().includes('该账户已被冻结')) {
        return { code: 400, msg: '该账户已被冻结' }
      }
      
      if ($('#errorcode').text().includes('验证码') || $('#errorpassword').text().includes('验证码')) {
        console.log(`验证码错误 (尝试 ${attempts + 1}/${maxAttempts})`)
        attempts++
        await new Promise(resolve => setTimeout(resolve, 1000)) // 延迟1秒重试
        continue
      }
      
      const isSuccess = $('.wrdvpn-navbar__title').text().trim() === '合肥工业大学WEBVPN系统' || 
                       $('.layui-show-sm-inline-block').text().trim() === '合肥工业大学WEBVPN系统'
      
      if (!isSuccess) {
        return { code: 400, msg: '账号或密码错误' }
      }
      
      success = true
      result = {
        code: 200,
        msg: '登录成功',
        data: {
          cookie: cookie1,
        },
      }
      
    } catch (err: any) {
      console.error('登录流程错误:', err)
      
      if (attempts < maxAttempts - 1) {
        attempts++
        await new Promise(resolve => setTimeout(resolve, 1000)) // 延迟1秒重试
      } else {
        return {
          code: 500,
          msg: '登录失败: ' + (err.message || '未知错误'),
          attempts
        }
      }
    }
  }

  if (!success) {
    return {
      code: 400,
      msg: '登录失败，验证码多次尝试错误',
      attempts
    }
  }

  return result
}