import App from './App';
import { createRoot } from 'react-dom/client';
import { BrowserRouter } from "react-router-dom";
import { Provider } from 'react-redux'
const container = document.getElementById('root');
const root = createRoot(container);

root.render(<App />);