import { Component, createEffect, createResource, createSignal, For } from 'solid-js';
import SingleItem from './components/Item';
import NewItemForm from './components/NewItemForm';
import { createItem, getItems } from './core/api';
import { Item, ItemCreate } from './core/types';

const App: Component = () => {
  const apiURL = (new URLSearchParams(document.location.search)).get("api")

  const [items, setItems] = createSignal<Item[]>([])
  const [resourceItems, { refetch: refetchItems }] = createResource(async () => {
    if (apiURL)
      return await getItems(apiURL)
    return []
  })

  createEffect(() => {
    // Loads existing items onto page
    let newItems = resourceItems()
    if (newItems)
      setItems(newItems)
  })

  const onNewItemSubmit = (newItem: ItemCreate) => {
    if (apiURL)
      createItem(apiURL, newItem).then((_) => {
        // TODO instead of refetching everything,
        // updating the existing page should be preferred
        refetchItems()
      })
  }
  const onItemDelete = (itemId: string) => {
    // TODO implement
  }

  return (
    <div class='md:container md:mx-auto'>
      <h1 class='text-center text-3xl'>FreeCycle</h1>
      <p>{apiURL}</p>
      <section class='mt-2'>
        <h2 class='text-center text-xl'>Create New</h2>
        <NewItemForm onSubmit={onNewItemSubmit} />
      </section>
      <section class='mt-2' data-page="items">
        <h2 class='text-center text-xl'>Items</h2>
        <ul>
          <For each={items()}>
            {item => <li><SingleItem item={item} onDeleteClick={onItemDelete} /></li>}
          </For>
        </ul>
      </section>
    </div>
  );
};

export default App;
