import { Component, createEffect, createResource, createSignal, For } from 'solid-js';
import SingleItem from './components/Item';
import NewItemForm from './components/NewItemForm';
import { createItem, deleteItem, getItems } from './core/api';
import { Item, ItemCreate } from './core/types';

const App: Component = () => {
  // SOURCE: https://developer.mozilla.org/en-US/docs/Web/API/URLSearchParams
  // SOURCE: https://stackoverflow.com/a/37832755
  const apiURL = (new URLSearchParams(document.location.search)).get("api")?.replace(/\/+$/, "")

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
    <>
      <header>
        <nav class="relative w-full flex flex-wrap items-center justify-between py-3 bg-gray-100 text-gray-500 hover:text-gray-700 focus:text-gray-700 shadow-lg">
          <div class="container-fluid w-full flex flex-wrap items-center justify-between px-6">
            <a class="text-3xl text-black" href="#">FreeCycle</a>
            <span class='truncate' style="max-width: 10rem;">{apiURL}</span>
          </div>
        </nav>
      </header>
      <main class='md:container md:mx-auto px-2'>
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
      </main>
    </>
  );
};

export default App;
