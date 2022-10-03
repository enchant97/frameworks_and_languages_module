import { Component, For } from 'solid-js';

export type ItemProps = {
  id: string
  userId: string
  keywords: string[]
  Description: string
  ImageUrl?: string
  Lat?: number
  Lon?: number
  DateFrom: Date
  onDeleteClick: (itemId: string) => void
}

const Item: Component<ItemProps> = (props) => {
  return (
    <>
      <span data-field="id">{props.id}</span>
      <img src={props.ImageUrl} data-field="image" />
      <span data-field="user_id">{props.userId}</span>
      <div>
        LatLon:
        <span data-field="lat">{props.Lat}</span>
        <span data-field="lon">{props.Lon}</span>
      </div>
      <time dateTime={props.DateFrom.toDateString()}>{props.DateFrom.toLocaleString()}</time>
      <p data-field="description">{props.Description}</p>
      <ul data-field="keywords">
        <For each={props.keywords}>
          {item => <li>{item}</li>}
        </For>
      </ul>
      <button data-action="delete">Delete</button>
    </>
  );
};

export default Item;
