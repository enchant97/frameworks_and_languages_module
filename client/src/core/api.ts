import { ItemCreate, Item } from "./types";

const DEFAULT_HEADERS = {
  'Content-Type': 'application/json',
}

export async function createItem(baseUrl: string, newItem: ItemCreate): Promise<Item> {
  let response = await fetch(
    baseUrl + '/item/',
    {
      method: 'POST',
      headers: DEFAULT_HEADERS,
      body: JSON.stringify(newItem),
    })
  if (!response.ok) {
    throw new Error(`request error ${response.status}`)
  }
  return await response.json()
}

export async function getItems(baseUrl: string): Promise<Item[]> {
  let response = await fetch(
    baseUrl + '/items/',
  )
  if (!response.ok) {
    throw new Error(`request error ${response.status}`)
  }
  return response.json()
}
