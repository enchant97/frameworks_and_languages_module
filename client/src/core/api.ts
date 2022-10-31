import { ItemCreate, Item } from "./types";

const DEFAULT_HEADERS = {
  'Content-Type': 'application/json',
}

/**
 * post a new item to api server
 */
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
  let data: Item = await response.json()
  // converts all the dates from strings into date
  data.date_from = new Date(data.date_from)
  if (data.date_to)
    data.date_to = new Date(data.date_to)
  return data
}

/**
 * Get all items from api server
 */
export async function getItems(baseUrl: string): Promise<Item[]> {
  let response = await fetch(
    baseUrl + '/items/',
  )
  if (!response.ok) {
    throw new Error(`request error ${response.status}`)
  }

  let data: Item[] = await response.json()

  // converts all the dates from strings into date
  data.forEach((item) => {
    item.date_from = new Date(item.date_from)
    if (item.date_to)
      item.date_to = new Date(item.date_to)
  })

  return data
}

/**
 * Delete an item by it's id from api server
 */
export async function deleteItem(baseUrl: string, itemId: string) {
  let response = await fetch(
    baseUrl + '/item/' + itemId + "/",
    {
      method: "DELETE",
    },
  )
  if (!response.ok) {
    throw new Error(`request error ${response.status}`)
  }
}
