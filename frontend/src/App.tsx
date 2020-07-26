import React from 'react';
import { Router, Route } from 'react-router-dom';
import './App.css';
import history from './history';
import Menu from './components/Menu';
import Login from "./components/Login";
import Signup from "./components/Signup";

function App() {
  return (
    <div className="App">
        <Router history={history}>
            <div>
                <Route exact path="/login" component={Login} />
                <Route exact path="/signup" component={Signup} />
                <Route exact path="/" component={Menu}/>
            </div>
        </Router>
    </div>
  );
}

export default App;
