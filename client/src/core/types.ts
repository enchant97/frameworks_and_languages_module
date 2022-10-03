/**
 * A created item,
 * received from the api server
 */
export type Item = {
  id: string
  user_id: string
  keywords: string[]
  description: string
  image?: string
  lat?: number
  lon?: number
  date_from: Date
  date_to?: Date
}

/**
 * Used to create a new item,
 * sending to the api server
 */
export type ItemCreate = {
  user_id: string
  keywords: string[]
  description: string
  image?: string
  lat?: number
  lon?: number
}
