import React from 'react';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom'
import './App.css';
import Menu from './components/Menu';
import Login from "./components/Login";

function App() {
  return (
    <div className="App">
        <Router>
            <Switch>
                <Route exact path="/login" component={Login} />
                <Route exact path="/" component={Menu}/>
            </Switch>
        </Router>
    </div>
  );
}

export default App;
