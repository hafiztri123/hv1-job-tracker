export type FetchResponse = {
  message: string
  status: number
}

export type FetchDetailResponse<T = string> = {
  message: string
  status: number
  data: T
}

export type FetchListResponse<T> = {
  message: string
  status: number
  data: {
    data: T[]
    dataCount: number
  }
}
