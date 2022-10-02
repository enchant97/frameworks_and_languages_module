import { Component } from 'solid-js';
import { Container } from 'solid-bootstrap';
import NewItemForm from './components/NewItemForm';


const App: Component = () => {
  let apiURL = (new URLSearchParams(document.location.search)).get("api")

  return (
    <Container>
      <h1>FreeCycle</h1>
      <p>{apiURL}</p>
      <section>
        <h2>Create New</h2>
        <NewItemForm />
      </section>
      <section data-page="items">
        <h2>Items</h2>
        <ul></ul>
      </section>
    </Container>
  );
};

export default App;
