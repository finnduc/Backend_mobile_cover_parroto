export interface CreateUserDto {
  email: string
  name: string
  password: string
}

export interface UpdateUserDto {
  name?: string
}
