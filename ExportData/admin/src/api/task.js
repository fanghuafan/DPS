import request from '@/utils/request'

export function get(params) {
  return request({
    url: '/api/v1/task/get',
    method: 'get',
    params
  })
}

export function add(params) {
  return request({
    url: '/api/v1/task/create',
    method: 'post',
    params
  })
}
