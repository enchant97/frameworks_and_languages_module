/**
 * A SolidJS component
 * allowing for reuse when displaying multiple items on screen
 */
import { Component, For } from 'solid-js';
import { Item } from '../core/types';

export type ItemProps = {
  item: Item
  onDeleteClick?: (itemId: string) => void
}

const SingleItem: Component<ItemProps> = (props) => {
  const onDelete = () => {
    props.onDeleteClick?.(props.item.id)
  }
  return (
    <>
      <span data-field="id">{props.item.id}</span>
      <img src={props.item.image} data-field="image" />
      <span data-field="user_id">{props.item.user_id}</span>
      <div>
        LatLon:
        <span data-field="lat">{props.item.lat}</span>
        <span data-field="lon">{props.item.lon}</span>
      </div>
      <time dateTime={props.item.date_from.toDateString()}>{props.item.date_from.toLocaleString()}</time>
      <p data-field="description">{props.item.description}</p>
      <ul data-field="keywords">
        <For each={props.item.keywords}>
          {item => <li>{item}</li>}
        </For>
      </ul>
      <button class='text-white bg-red-500 hover:bg-red-800 rounded p-2' data-action="delete" onclick={onDelete}>Delete</button>
    </>
  );
};

export default SingleItem;
