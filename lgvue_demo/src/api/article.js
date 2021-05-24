import request from '@/utils/request'

export function createarticle (data) {
  return request({
    url: '/lego/article/createarticle',
    method: 'post',
    data
  })
}
export function deletearticle (data) {
  return request({
    url: '/lego/article/deletearticle',
    method: 'post',
    data
  })
}

export function getarticle (data) {
  return request({
    url: '/lego/article/getarticle',
    method: 'post',
    data
  })
}

export function getauthoIrdarticles (data) {
  return request({
    url: '/lego/article/getauthoIrdarticles',
    method: 'post',
    data
  })
}

