import { Component } from 'solid-js';
import { Form, Button } from 'solid-bootstrap';

const NewItemForm: Component = () => {
    return (
      <Form>
        <Form.Group class="mb-3">
          <Form.Label>User ID</Form.Label>
          <Form.Control type='text' name='user_id' required />
        </Form.Group>

        <Form.Group class="mb-3">
          <Form.Label>Lat</Form.Label>
          <Form.Control type='number' name='lat' />
        </Form.Group>

        <Form.Group class="mb-3">
          <Form.Label>Lon</Form.Label>
          <Form.Control type='number' name='lon' />
        </Form.Group>

        <Form.Group class="mb-3">
          <Form.Label>Description</Form.Label>
          <Form.Control as='textarea' name='description' rows={3} required />
        </Form.Group>

        <Form.Group class="mb-3">
          <Form.Label>Image</Form.Label>
          <Form.Control type='url' name='image' />
        </Form.Group>

        <Form.Group class="mb-3">
          <Form.Label>Keywords</Form.Label>
          <Form.Control type='text' name='keywords' />
        </Form.Group>

        <Button variant='primary' type='submit' data-action='create_item'>Submit</Button>
      </Form>
    );
  };


export default NewItemForm;
