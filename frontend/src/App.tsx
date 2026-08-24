import { lazy, onMount, type ParentProps } from 'solid-js';
import { Router, Route } from "@solidjs/router";
import './App.css';
import { runAtStart } from './functions/atstart';

import Body from './pages/Body';
import Header from './components/Header';

function AppLayout(props: ParentProps) {
  return (
    <>
      <Header></Header>
      <div class="container-lg">
        <div class="row">
          <div class="col-md mt-4 mb-4">
            {props.children}
          </div>
        </div>
      </div>
    </>
  )
}

function App() {

  onMount(() => {
    runAtStart();
  });

  const Config = lazy(() => import("./pages/Config"));
  const History = lazy(() => import("./pages/History"));
  const Activity = lazy(() => import("./pages/Activity"));
  const HostPage = lazy(() => import("./pages/HostPage"));

  return (
    <Router root={AppLayout}>
      <Route path="/" component={Body}/>
      <Route path="/config" component={Config}/>
      <Route path="/history" component={History}/>
      <Route path="/activity" component={Activity}/>
      <Route path="/host/:id" component={HostPage}/>
    </Router>
  )
}

export default App
