import React from 'react';
import Home from './components/Home';
import {Employee, AddEmployee, UpdateEmployee, DeleteEmployee} from './components/Employee';
import {BrowserRouter as Router, Route,Routes} from 'react-router-dom';

function App() {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/employees" element={<Employee />} />
          <Route path="/add-employee" element={<AddEmployee />} />
        
          <Route path="/update-employee/:id" element={<UpdateEmployee />} />
          <Route path="/delete-employee" element={<DeleteEmployee />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
