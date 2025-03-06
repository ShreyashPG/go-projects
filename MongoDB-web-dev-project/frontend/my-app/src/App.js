
import React from 'react';
import {Employees, CreateEmployee, GetEmployeeByID, UpdateEmployee, DeleteEmployee }from './components/Employee';
import {BrowserRouter as Router, Route, Routes} from 'react-router-dom';




function App() {

  return (
    <div className="App">
      <Router>
        <Routes>
          <Route path="/" element={<Employees />} />
          <Route path="/create-employee" element={<CreateEmployee />} />
          <Route path="/employee/:id" element={<GetEmployeeByID />} />
          <Route path="/update-employee/:id" element={<UpdateEmployee />} />
          <Route path="/delete-employee" element={<DeleteEmployee />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
