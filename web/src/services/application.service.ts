import { createAxiosInstance } from '@/utils/createAxiosInstance'
import type { AxiosResponse } from 'axios'
import type { FetchDetailResponse, FetchListResponse } from './type/response.type'
import type { CreateApplicationDto, UpdateApplicationDto, Application } from './dto/application.dto'

const API = createAxiosInstance('applications')

const ApplicationServices = {
  getApplications: (status?: string, limit?: number, offset?: number): Promise<AxiosResponse<FetchListResponse<Application>>> => {
    const params: Record<string, string | number> = {}
    if (status) params.status = status
    if (limit !== undefined) params.limit = limit
    if (offset !== undefined) params.offset = offset
    return API.get('/', { params })
  },
  createApplication: (body: CreateApplicationDto): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.post('/', body)
  },
  updateApplication: (id: string, body: UpdateApplicationDto): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.put(`/${id}`, body)
  },
  deleteApplication: (id: string): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.delete(`/${id}`)
  },
  getApplicationOptions: (): Promise<AxiosResponse<FetchDetailResponse<{ statusOption: string[] }>>> => {
    return API.get('/options', { params: { statusOption: true } })
  },
  batchDeleteApplications: (applicationIds: string[]): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.delete('/batch/delete', { data: { applicationIds } })
  },
  batchUpdateStatusApplications: (applicationIds: string[], status: string): Promise<AxiosResponse<FetchDetailResponse>> => {
    return API.put('/batch/status', { applicationIds, status })
  }
}

export default ApplicationServices
