import { Component } from 'solid-js';

const NewItemForm: Component = () => {
    return (
      <form class='max-w-lg mx-auto flex flex-col'>
        <div class='grid grid-cols-2 mb-2'>
          <label>User ID</label>
          <input class='border-2 rounded' type='text' name='user_id' required />
          <label>Lat</label>
          <input class='border-2 rounded' type='number' name='lat' />
          <label>Lon</label>
          <input class='border-2 rounded' type='number' name='lon' />
          <label>Description</label>
          <textarea class='border-2 rounded' name="description" rows="3" required></textarea>
          <label>Image</label>
          <input class='border-2 rounded' type='url' name='image' />
          <label>Keywords</label>
          <input class='border-2 rounded' type='text' name='keywords' />
        </div>
        <button class='text-white bg-sky-500 hover:bg-sky-800 rounded p-2' type='submit' data-action='create_item'>Submit</button>
      </form>
    );
  };


export default NewItemForm;
