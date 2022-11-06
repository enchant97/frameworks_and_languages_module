/* @refresh reload */
import { render } from 'solid-js/web';

import './index.css';
import App from './App';

// SOURCE: https://www.solidjs.com/tutorial/introduction_basics
render(() => <App />, document.getElementById('root') as HTMLElement);
