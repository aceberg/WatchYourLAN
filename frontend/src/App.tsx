import { lazy, onMount } from 'solid-js';
import { Router, Route } from "@solidjs/router";
import './App.css';
import { runAtStart } from './functions/atstart';

import Body from './pages/Body';
import Header from './components/Header';

function App() {

  onMount(() => {
    runAtStart();
  });

  const Config = lazy(() => import("./pages/Config"));
  const History = lazy(() => import("./pages/History"));
  const HostPage = lazy(() => import("./pages/HostPage"));

  return (
    <>
    <Header></Header>
    <div class="container-lg">
      <div class="row">
        <div class="col-md mt-4 mb-4">
          <Router>
            <Route path="/" component={Body}/>
            <Route path="/config" component={Config}/>
            <Route path="/history" component={History}/>
            <Route path="/host/:id" component={HostPage}/>
          </Router>
        </div>
      </div>
    </div>
    </>
  )
}

export default App
