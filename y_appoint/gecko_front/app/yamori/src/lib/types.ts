export interface LoginSession {
  user: {
    email: string,
    accessToken?: string
  }
}