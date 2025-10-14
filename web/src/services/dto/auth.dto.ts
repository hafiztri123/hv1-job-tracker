export type LoginBody = {
  email: string
  password: string
}

export type RegisterBody = LoginBody & {
  firstName: string
  lastName: string
}
