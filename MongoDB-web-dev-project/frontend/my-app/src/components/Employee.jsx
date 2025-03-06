import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { useParams } from 'react-router-dom';
import { useNavigate } from 'react-router-dom';


export const Employees = () => {
    const [employees, setEmployees] = useState([]);

    useEffect(()=>{
        axios.get('http://localhost:3001/')
        .then(response => {
            // setEmployees(response.data);
            setEmployees(Array.isArray(response.data.data) ? response.data.data : []);
            console.log(response.data);
        })
        .catch(error => {
            console.error('Error fetching employees:', error);
        });
    }, []);

    return (
        <div>
            <h1>Employees</h1>
            <ul>
                {employees.map(employee => (
                    <li key={employee.Id}>{employee.Name} - {employee.Email} - {employee.Password}</li>
                ))}
            </ul>   
        </div>

    );
};

export const GetEmployeeByID = () => {
    const [employee, setEmployee] = useState({
        Name: '',
        Email: '',
        Password: ''
    });
    const {id} = useParams();


    useEffect(()=>{
        axios.get(`http://localhost:3001/${id}`)
        .then(response => {
            console.log(response.data);
            setEmployee(response.data.data);
        })
        .catch(error => {
            console.error('Error fetching employee:', error);
        });
    }, [id]);
    return (
        <div>
            <h1>Employee</h1>
            <p>Name: {employee.Name}</p>
            <p>Email: {employee.Email}</p>
            <p>Phone: {employee.Password}</p>
         
        </div>
    );
};

export const CreateEmployee =( )=>{
    const [employee, setEmployee] = useState({
        Name: '',
        Email: '',
        Password: ''
    });
    // const navigate = useNavigate();
    const handleSubmit = (e) => {
        e.preventDefault();
        if (!employee.Name || !employee.Email || !employee.Password) {
            alert("All fields are required!");
            return;
        }
        axios.post('http://localhost:3001/', employee)
        .then(
            () => alert("Employee created successfully!"),
            console.log(employee),
            // flush the form
            setEmployee({
                Name: '',
                Email: '',
                Password: ''
            }),
            // //redirect to the employees page
            // navigate('/')
        )
        .catch(error => console.error('Error creating employee:', error));
    };
    
    return (
        <div>
            <h1>Create Employee</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" name="name" value={employee.Name} onChange={(e)=>setEmployee({...employee, Name: e.target.value})} />
                <input type="text" name="email" value={employee.Email} onChange={(e)=>setEmployee({...employee, Email: e.target.value})} />
                <input type="text" name="password" value={employee.Password} onChange={(e)=>setEmployee({...employee, Password: e.target.value})} />
                <button type="submit">Create Employee</button>
            </form>
        </div>
    )
    
}

export const UpdateEmployee = () => {
    const [employee, setEmployee] = useState({
        Name: '',
        Email: '',
        Password: ''
    });
    const navigate = useNavigate();
    const {id} = useParams();
    useEffect(()=>{
        axios.get(`http://localhost:3001/${id}`)
        .then(response => {
            setEmployee(response.data);
            alert("Employee fetched successfully!");
            console.log(response.data);
            navigate('/');
        })
        .catch(error => {
            console.error('Error fetching employee:', error);
        });
    }, [id]);

    const handleSubmit = (e) => {
        e.preventDefault();
        axios.put(`http://localhost:3001/${id}`, employee);
    };
    return (
        <div>
            <h1>Update Employee</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" name="name" value={employee.Name} onChange={(e)=>setEmployee({...employee, Name: e.target.value})} />
                <input type="text" name="email" value={employee.Email} onChange={(e)=>setEmployee({...employee, Email: e.target.value})} />
                <input type="text" name="password" value={employee.Password} onChange={(e)=>setEmployee({...employee, Password: e.target.value})} />
                <button type="submit">Update Employee</button>
            </form>
        </div>
    )
}

export const DeleteEmployee = () => {
    const [id, setId] = useState('');
    const navigate = useNavigate();
    const handleSubmit = (e) => {
        e.preventDefault();
        if (window.confirm("Are you sure you want to delete this employee?")) {
            axios.delete(`http://localhost:3001/${id}`)
            .then(() => alert("Employee deleted successfully!") ,
                navigate('/')
        )
            .catch(error => console.error('Error deleting employee:', error));
        }
    };
    
    return (
        <div>
            <h1>Delete Employee</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" name="id" value={id} onChange={(e)=>setId(e.target.value)} />
                <button type="submit">Delete Employee</button>
            </form>
        </div>
    )
}

