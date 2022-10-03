import { Component } from 'solid-js';
import NewItemForm from './components/NewItemForm';


const App: Component = () => {
  let apiURL = (new URLSearchParams(document.location.search)).get("api")

  return (
    <div class='md:container md:mx-auto'>
      <h1 class='text-center text-3xl'>FreeCycle</h1>
      <p>{apiURL}</p>
      <section class='mt-2'>
        <h2 class='text-center text-xl'>Create New</h2>
        <NewItemForm />
      </section>
      <section class='mt-2' data-page="items">
        <h2 class='text-center text-xl'>Items</h2>
        <ul></ul>
      </section>
    </div>
  );
};

export default App;
