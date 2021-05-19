import request from '@/utils/request'

export function registerbycaptcha (data) {
  return request({
    url: '/lego/user/registerbycaptcha',
    method: 'post',
    data
  })
}

export function loginbycaptcha (data) {
  return request({
    url: '/lego/user/loginbycaptcha',
    method: 'post',
    data
  })
}

export function loginbypassword (data) {
  return request({
    url: '/lego/user/loginbypassword',
    method: 'post',
    data
  })
}

export function getuserinfo (data) {
  return request({
    url: '/lego/user/getuserinfo',
    method: 'post',
    data
  })
}


export function loginout (data) {
  return request({
    url: '/lego/user/loginout',
    method: 'post',
    data
  })
}
