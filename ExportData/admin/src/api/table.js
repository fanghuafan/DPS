import request from '@/utils/request'

export function getList(params) {
  return request({
    url: '/api/v1/list/list',
    method: 'get',
    params
  })
}
