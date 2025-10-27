import { createAxiosInstance } from '@/utils/createAxiosInstance'
import type { AxiosResponse } from 'axios'
import type { FetchDetailResponse } from './type/response.type'
import type { LoginBody, RegisterBody } from './dto/auth.dto'

const API = createAxiosInstance('auth')

const AuthServices = {
  login: (body: LoginBody): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.post('/login', body)
  },
  register: (body: RegisterBody): Promise<AxiosResponse> => {
    return API.post('/register', body)
  },
  verify: (): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.get('/verify')
  },
  logout: (): Promise<AxiosResponse> => {
    return API.post('/logout')
  }
}

export default AuthServices
