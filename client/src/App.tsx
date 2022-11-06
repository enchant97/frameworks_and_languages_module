import { Component, createEffect, createResource, createSignal, For } from 'solid-js';
import SingleItem from './components/Item';
import NewItemForm from './components/NewItemForm';
import { createItem, deleteItem, getItems } from './core/api';
import { Item, ItemCreate } from './core/types';

const App: Component = () => {
  // SOURCE: https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams
  const apiURL = (new URLSearchParams(document.location.search)).get("api")

  // SOURCE: https://www.solidjs.com/tutorial/introduction_signals
  const [items, setItems] = createSignal<Item[]>([])
  // SOURCE: https://www.solidjs.com/tutorial/async_resources
  const [resourceItems, { refetch: refetchItems }] = createResource(async () => {
    if (apiURL)
      return await getItems(apiURL)
    // if no api url has been set, just act like there are no items to display
    return []
  })

  // SOURCE: https://www.solidjs.com/tutorial/introduction_effects
  createEffect(() => {
    // Loads existing items onto page
    let newItems = resourceItems()
    if (newItems)
      setItems(newItems)
  })

  const onNewItemSubmit = (newItem: ItemCreate) => {
    if (apiURL)
      createItem(apiURL, newItem).then(_ => {
        // TODO instead of refetching everything,
        // updating the existing page should be preferred
        refetchItems()
      })
  }
  const onItemDelete = (itemId: string) => {
    if (apiURL)
      deleteItem(apiURL, itemId).then(_ => {
        // TODO instead of refetching everything,
        // updating the existing page should be preferred
        refetchItems()
      })
  }

  return (
    <div class='md:container md:mx-auto px-2'>
      <h1 class='text-center text-3xl'>FreeCycle</h1>
      <p class='text-center'>{apiURL}</p>
      <section class='mt-2'>
        <h2 class='text-center text-xl'>Create New</h2>
        <NewItemForm onSubmit={onNewItemSubmit} />
      </section>
      <section class='mt-2' data-page="items">
        <h2 class='text-center text-xl'>Items</h2>
        <ul class='grid grid-cols-3 gap-2'>
          <For each={items()}>
            {item => <li class='border-2 rounded p-2'><SingleItem item={item} onDeleteClick={onItemDelete} /></li>}
          </For>
        </ul>
      </section>
    </div>
  );
};

export default App;
