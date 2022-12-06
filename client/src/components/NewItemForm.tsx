/**
 * A SolidJS component
 * Separating out logic from the App component
 */
import { Component } from 'solid-js';
import { createStore } from 'solid-js/store';
import { ItemCreate } from '../core/types';

type NewItemFormProps = {
  onSubmit?: (item: ItemCreate) => void
}

const NewItemForm: Component<NewItemFormProps> = (props) => {
  // SOURCE: https://www.solidjs.com/tutorial/stores_createstore
  // Using a store to group related signals into one
  const [fields, setFields] = createStore<ItemCreate>({
    user_id: "",
    keywords: [],
    description: "",
  })

  const handleOnSubmit = (e: SubmitEvent) => {
    e.preventDefault();
    // TODO add validation?
    props.onSubmit?.(fields)
    // reset form to defaults
    setFields({
      user_id: "",
      keywords: [],
      description: "",
      image: undefined,
      lat: undefined,
      lon: undefined,
    })
  }

  return (
    <form class='max-w-lg mx-auto flex flex-col' onsubmit={handleOnSubmit}>
      <div class='grid grid-cols-2 mb-2'>
        <label>User ID</label>
        <input
          class='border-2 rounded' type='text' name='user_id' required
          value={fields.user_id}
          oninput={(e) => setFields('user_id', e.target.value)}
        />
        <label>Lat</label>
        <input
          class='border-2 rounded' type='number' name='lat'
          value={fields.lat || 0}
          oninput={(e) => setFields('lat', parseFloat(e.target.value))}
        />
        <label>Lon</label>
        <input
          class='border-2 rounded' type='number' name='lon'
          value={fields.lon || 0}
          oninput={(e) => setFields('lon', parseFloat(e.target.value))}
        />
        <label>Description</label>
        <textarea
          class='border-2 rounded' name="description" rows="3" required
          value={fields.description}
          oninput={(e) => setFields('description', e.target.value)}
        ></textarea>
        <label>Image</label>
        <input
          class='border-2 rounded' type='url' name='image'
          value={fields.image || ""}
          oninput={(e) => setFields('image', e.target.value)}
        />
        <label>Keywords</label>
        <input
          class='border-2 rounded' type='text' name='keywords'
          value={fields.keywords.join(",")}
          oninput={(e) => setFields('keywords', e.target.value.split(","))}
        />
      </div>
      <button
        class='text-white bg-sky-500 hover:bg-sky-800 rounded p-2'
        type='submit' data-action='create_item'
      >Submit</button>
    </form>
  );
};


export default NewItemForm;
