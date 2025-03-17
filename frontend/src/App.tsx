import { onMount } from 'solid-js';
import './App.css'
import { apiGetAllHosts } from './functions/api';
import { setAllHosts } from './functions/exports';
import Body from './pages/Body'

function App() {

  onMount(async () => {

    const hosts = await apiGetAllHosts();
    setAllHosts(hosts);
  });

  return (
    <div class="container-lg">
      <div class="row">
        <div class="col-md mt-4 mb-4">
          <Body></Body>
        </div>
      </div>
    </div>
  )
}

export default App
