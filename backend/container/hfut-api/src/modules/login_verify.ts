import * as CryptoJS from 'crypto-js'
import { AxiosError } from 'axios'
import * as cheerio from 'cheerio'
import request, { getCookie } from '../shared/request'
import { IQuery } from '../server'
import Tesseract from 'tesseract.js'
import * as fs from 'fs'
import * as path from 'path'

// 验证码识别函数
async function recognizeCaptcha(imageBuffer: Buffer): Promise<string> {
  console.log('[验证码识别] 开始识别验证码')
  
  try {
    const worker = await Tesseract.createWorker('eng')
    await worker.setParameters({
      tessedit_char_whitelist: 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
    })

    console.log('[验证码识别] OCR引擎初始化完成')
    
    // 保存验证码图片用于调试
    const debugDir = path.join(__dirname, '..', 'debug')
    if (!fs.existsSync(debugDir)) {
      fs.mkdirSync(debugDir, { recursive: true })
    }
    const imagePath = path.join(debugDir, `captcha-${Date.now()}.png`)
    fs.writeFileSync(imagePath, imageBuffer)
    console.log(`[验证码识别] 验证码图片已保存至: ${imagePath}`)

    const { data: { text } } = await worker.recognize(imageBuffer)
    await worker.terminate()
    
    const cleanedText = text.replace(/\s+|\W+/g, '').slice(0, 4) || 'ABCD'
    
    console.log(`[验证码识别] OCR原始识别结果: ${text}`)
    console.log(`[验证码识别] 清理后验证码: ${cleanedText}`)
    
    return cleanedText
  } catch (error) {
    console.error('[验证码识别] OCR识别错误:', error)
    return 'ABCD' // 默认值防止空结果
  }
}

const url1 = 'https://cas.hfut.edu.cn/cas/login?service=https%3A%2F%2Fcas.hfut.edu.cn%2Fcas%2Foauth2.0%2FcallbackAuthorize%3Fclient_id%3DBsHfutEduPortal%26redirect_uri%3Dhttps%253A%252F%252Fone.hfut.edu.cn%252Fhome%252Findex%26response_type%3Dcode%26client_name%3DCasOAuthClient'

function encryptionPwd(pwd: string, salt: string) {
  let key = CryptoJS.enc.Utf8.parse(salt)
  let password = CryptoJS.enc.Utf8.parse(pwd)
  let encrypted = CryptoJS.AES.encrypt(password, key, { 
    mode: CryptoJS.mode.ECB, 
    padding: CryptoJS.pad.Pkcs7 
  })
  return encrypted.toString()
}

export default async function(query: IQuery, getPwd = false) {
  // 首先解构出 req 和 res
  const { req, res } = query
  
  // 现在可以安全地访问 req
  console.log('[登录流程] 开始登录流程')
  console.log(`[登录流程] 用户名: ${req.query.username}`)
  
  const username = req.query.username as string
  const password = req.query.password as string

  if (!username || !password) {
    console.log('[登录流程] 用户名或密码为空')
    return {
      code: 400,
      msg: '用户名或密码不能为空',
    }
  }

  const maxAttempts = 3
  let attempts = 0
  let success = false
  let result = null
  let payload = { cookie: '' }

  while (attempts < maxAttempts && !success) {
    console.log(`[登录流程] 尝试 #${attempts + 1}/${maxAttempts}`)
    
    try {
      console.log('[登录流程] 第一步: 获取初始页面')
      const res1 = await request(url1, {}, query)
      let cookie1 = getCookie(res1.cookie as string[])
      payload = { cookie: cookie1 }

      if (!cookie1) {
        console.log('[登录流程] 无法获取初始Cookie')
        return {
          code: 400,
          msg: '登录太过频繁，请稍后再试'
        }
      }

      console.log(`[登录流程] 初始Cookie: ${cookie1}`)
      
      console.log('[登录流程] 第二步: 获取验证码图片')
      const vercodeRes = await request('https://cas.hfut.edu.cn/cas/vercode', {
        responseType: 'arraybuffer'
      }, payload)

      const imageBuffer = vercodeRes.body
      const capchaText = await recognizeCaptcha(imageBuffer)

      console.log(`[登录流程] 识别出的验证码: ${capchaText}`)

      if (vercodeRes.cookie && vercodeRes.cookie.length > 0) {
        payload.cookie += `; ${vercodeRes.cookie[0]}`
        console.log(`[登录流程] 更新Cookie (验证码): ${vercodeRes.cookie[0]}`)
      }

      console.log('[登录流程] 第三步: 检查验证码状态')
      const url2 = `https://cas.hfut.edu.cn/cas/checkInitVercode?_=${Date.now()}`
      const res2 = await request(url2, {}, payload)
      
      if (res2.cookie && res2.cookie.length > 0) {
        payload.cookie += `; ${res2.cookie[0]}`
        console.log(`[登录流程] 更新Cookie (验证码状态): ${res2.cookie[0]}`)
      }
      
      // 清理Cookie格式
      payload.cookie = payload.cookie.replace(' Path=/cas/; Secure; HttpOnly;', '')
      console.log(`[登录流程] 清理后Cookie: ${payload.cookie}`)

      // 安全获取加密盐
      const saltKey = res2.cookie?.[0]?.split('=')[1] || ''
      console.log(`[登录流程] 获取的加密盐值: ${saltKey}`)

      if (!saltKey) {
        console.log('[登录流程] 无法获取加密盐值')
        return {
          code: 500,
          msg: '无法获取加密盐值'
        }
      }

      const encryptedPwd = encryptionPwd(password, saltKey)
      console.log(`[登录流程] 加密后的密码: ${encryptedPwd}`)

      console.log('[登录流程] 第四步: 验证用户身份')
      // 使用URLSearchParams安全构建查询参数
      const checkParams = new URLSearchParams()
      checkParams.append('username', username)
      checkParams.append('password', encryptedPwd)
      checkParams.append('capcha', capchaText)
      checkParams.append('_', Date.now().toString())
      
      const url3 = `https://cas.hfut.edu.cn/cas/policy/checkUserIdenty?${checkParams.toString()}`
      console.log(`[登录流程] 请求URL: ${url3}`)
      
      const res3 = await request(url3, {}, payload)
      console.log(`[登录流程] 身份验证响应: ${JSON.stringify(res3.body)}`)

      console.log('[登录流程] 第五步: 提交登录请求')
      const loginParams = new URLSearchParams({
        username,
        capcha: capchaText,
        execution: 'e1s1',
        _eventId: 'submit',
        password: encryptedPwd,
        geolocation: '',
      }).toString()
      
      console.log(`[登录流程] 登录参数: ${loginParams}`)
      
      // 发送登录请求（禁止自动重定向）
      await request(url1, {
        url: url1,
        method: 'POST',
        maxRedirects: 0, // 禁止自动重定向
        data: loginParams,
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded',
          'Cookie': payload.cookie
        }
      }, payload)

      console.warn('[登录流程] 登录请求未抛出异常，可能成功但没有重定向')
      result = {
        code: 200,
        msg: '登录成功',
        data: {
          ticketCode: 'unknown - no redirection',
          cookie: payload.cookie
        }
      }
      success = true
      
    } catch (err: any) {
      console.error('[登录流程] 捕获到异常:', err)
      
      if (err.isAxiosError) {
        const axiosError = err as AxiosError
        
        // 处理302重定向（登录成功）
        if (axiosError.response?.status === 302) {
          const location = axiosError.response.headers?.['location']
          console.log(`[登录流程] 重定向位置: ${location}`)
          
          if (location?.includes('ticket=')) {
            const ticketCode = location.split('ticket=')[1]
            console.log(`[登录流程] 提取到的ticket: ${ticketCode}`)
            
            // 更新cookie
            const cookies = axiosError.response.headers['set-cookie'] || []
            if (cookies.length > 0) {
              const newCookie = cookies[0].split(';')[0]
              payload.cookie += `; ${newCookie}`
              console.log(`[登录流程] 更新后的Cookie: ${newCookie}`)
            }

            result = {
              code: 200,
              msg: '登录成功',
              data: {
                ticketCode,
                cookie: payload.cookie,
              },
            }
            success = true
            break
          }
        }
        
        let errMsg = ''
        if (axiosError.response?.data) {
          console.log('[登录流程] 尝试解析错误页面')
          
          try {
            // 尝试解析HTML错误页面
            if (typeof axiosError.response.data === 'string') {
              const $ = cheerio.load(axiosError.response.data)
              errMsg = $('#errorpassword').text().trim() || 
                       $('#errorcode').text().trim() || 
                       $('.alert-error').text().trim() ||
                       ''
            } 
            // 尝试解析JSON错误响应
            else if (typeof axiosError.response.data === 'object') {
              errMsg = axiosError.response.data.message || 
                       axiosError.response.data.error ||
                       ''
            }
            
            if (errMsg) {
              console.log(`[登录流程] 从页面提取的错误信息: ${errMsg}`)
            }
          } catch (parseError) {
            console.error('[登录流程] 错误解析错误:', parseError)
          }
        }
        
        if (!errMsg) {
          errMsg = '未知登录错误'
        }
        
        // 检查是否是验证码错误
        if (errMsg.includes('验证码') || errMsg.includes('不正确')) {
          console.log(`[登录流程] 验证码错误 (尝试 ${attempts + 1}/${maxAttempts}): ${errMsg}`)
          attempts++
          // 添加延时避免过快请求
          await new Promise(resolve => setTimeout(resolve, 1000))
          continue
        }
        
        result = {
          code: 400,
          msg: errMsg,
        }
        attempts = maxAttempts // 非验证码错误，立即退出
        
      } else {
        // 非Axios错误
        console.error('[登录流程] 非Axios错误:', err)
        result = {
          code: 500,
          msg: `服务器错误: ${err.message}`,
        }
        attempts = maxAttempts // 立即退出
      }
    }
  }

  if (!success) {
    console.log(`[登录流程] 登录失败: ${result?.msg || '未知原因'}`)
    return result || {
      code: 400,
      msg: '登录失败，未知错误'
    }
  }

  console.log('[登录流程] 登录成功')
  return result
}