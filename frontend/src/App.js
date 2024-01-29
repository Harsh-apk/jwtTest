import './App.css';
import { BrowserRouter as Router, Route, Switch } from 'react-router-dom';
import Navbar from './components/Navbar';
import { useState } from 'react';
import LoginCreate from './components/LoginCreate';
import Home from './components/Home';
import Signup from './components/Signup';
import Login from './components/Login';

function App() {
  const [user,setUser] = useState(null);
  return (
      <Router>
        <div className="App" >
          <Navbar/>
          <Switch>
            <Route path="/" exact >
              <Home user={user} setUser={setUser} />
            </Route>
            <Route path="/login" ><Login user={user} setUser={setUser}/></Route>
            <Route path="/signup" ><Signup/></Route>
            <Route path="/loginCreate" ><LoginCreate/></Route>


          </Switch>




        </div>

      </Router>
  );
}

export default App;
