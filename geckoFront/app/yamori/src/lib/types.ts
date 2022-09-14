export interface Session {
  user: {
    email: string,
    accessToken?: string
  } | null
}