import React from 'react';
import Home from './components/Home';

import {BrowserRouter as Router, Route,Routes} from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          {/* <Route path="/employees" element={<Employee />} />
          <Route path="/add-employee/:id" element={<Employee />} />
          <Route path="/update-employee/:id" element={<Employee />} />
          <Route path="/delete-employee/:id" element={<Employee />} /> */}
        </Routes>
      </Router>
    </div>
  );
}

export default App;
