export type CreateApplicationDto = {
  companyName: string
  positionTitle: string
  jobUrl?: string
  salaryRange?: string
  location?: string
  status?: string
  notes?: string
  appliedDate?: string
}

export type UpdateApplicationDto = {
  companyName?: string
  positionTitle?: string
  jobUrl?: string
  salaryRange?: string
  location?: string
  status?: string
  notes?: string
  appliedDate?: string
}

export type Application = {
  id: string
  userId: string
  companyName: string
  positionTitle: string
  jobUrl?: string
  salaryRange?: string
  location?: string
  status?: string
  notes?: string
  appliedDate?: string
  createdAt: string
  updatedAt?: string
  deletedAt?: string
}
