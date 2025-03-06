import React, { useEffect, useState } from 'react';
import axios from 'axios';

const Employee =()=>{
    const [employees, setEmployees]=useState([]);
    const [loading, setLoading]=useState(true);
    useEffect(()=>{
        axios.get('http://localhost:3001/employees')
        .then((response)=>{
            setEmployees(response.data);
            setLoading(false);
        })
        .catch((error)=>{
            console.error('Error fetching employees:', error);
        });
    },[]);

    if(loading){
        return <div>Loading...</div>;
    }

    return(
        <div>
            <h1>Employee List</h1>
            <ul>
                {employees.map((employee)=>(
                    <li key={employee.id}>
                        {employee.name} - {employee.email} - {employee.age} - {employee.position}
                    </li>
                ))}
            </ul>
        </div>
    )
}


const AddEmployee=()=>{
    const [employee, setEmployee]=useState({
        name:'',
        email:'',
        age:'',
        position:''
    });


    // useEffect(()=>{
    //     axios.post('http://localhost:3001/add-employee', employee)
    //     .then((response)=>{
    //         console.log(response.data);
    //     })
    //     .catch((error)=>{
    //         console.error('Error adding employee:', error);
    //     })
    // },[employee]);
    function handleSubmit(e) {
        e.preventDefault();
        
        // Convert age to a number before sending
        const newEmployee = { ...employee, age: Number(employee.age) };
    
        axios.post('http://localhost:3001/add-employee', newEmployee, {
            headers: {
                'Content-Type': 'application/json'
            }
        })
            .then((response) => {
                alert('Employee added successfully');
                //flush the form 
                setEmployee({
                    name: '',
                    email: '',
                    age: '',
                    position: ''
                });
                console.log(response.data);
            })
            .catch((error) => {
                console.error('Error adding employee:', error.response?.data || error.message);
            });
    }
    
    return (
        <div>
            <h1>Add Employee</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" placeholder="Name" value={employee.name} onChange={(e)=>setEmployee({...employee, name:e.target.value})}/>
                <input type="email" placeholder="Email" value={employee.email} onChange={(e)=>setEmployee({...employee, email:e.target.value})}/>
                <input type="number" placeholder="Age" value={employee.age} onChange={(e)=>setEmployee({...employee, age:e.target.value})}/>
                <input type="text" placeholder="Position" value={employee.position} onChange={(e)=>setEmployee({...employee, position:e.target.value})}/>
                <button type="submit">Add Employee</button>
            </form>
        </div>
    )
    
}

const UpdateEmployee = () => {
    const [employee, setEmployee] = useState({
        name: '',
        email: '',
        age: '',
        position: ''
    });
    const [loading, setLoading] = useState(true);
    const [id, setId] = useState(null);

    useEffect(() => {
        const idFromURL = window.location.pathname.split('/').pop();
        setId(idFromURL);

        // Fetch existing employee data
        axios.get(`http://localhost:3001/employees/${idFromURL}`)
            .then((response) => {
                setEmployee(response.data); // Set fetched data
                setLoading(false);
            })
            .catch((error) => {
                console.error("Error fetching employee:", error.response?.data || error.message);
                setLoading(false);
            });
    }, []);

    function handleSubmit(e) {
        e.preventDefault();
        axios.put(`http://localhost:3001/update-employee/${id}`, employee)
            .then((response) => {
                console.log(response.data);
                alert("Employee updated successfully!");
            })
            .catch((error) => {
                console.error("Error updating employee:", error.response?.data || error.message);
            });
    }

    if (loading) {
        return <div>Loading...</div>;
    }

    return (
        <div>
            <h1>Update Employee</h1>
            <form onSubmit={handleSubmit}>
                <input type="text" placeholder="Name" value={employee.name} onChange={(e) => setEmployee({ ...employee, name: e.target.value })} />
                <input type="email" placeholder="Email" value={employee.email} onChange={(e) => setEmployee({ ...employee, email: e.target.value })} />
                <input type="number" placeholder="Age" value={employee.age} onChange={(e) => setEmployee({ ...employee, age: e.target.value })} />
                <input type="text" placeholder="Position" value={employee.position} onChange={(e) => setEmployee({ ...employee, position: e.target.value })} />
                <button type="submit">Update Employee</button>
            </form>
        </div>
    );
};


const DeleteEmployee=()=>{
    const [id, setId]=useState(null);
   function handleDelete(){
    axios.delete(`http://localhost:3001/delete-employee/${id}`)
    .then((response)=>{
        console.log(response.data);
        alert("Employee deleted successfully!");
    })
   }
    return(
        <div> 
            <h1>Delete Employee</h1>
            <p>Enter the id of Employee to delete</p>
            <input type="number" placeholder="ID" value={id} onChange={(e)=>setId(e.target.value)}/>
            <button onClick={handleDelete}>Delete</button>
        </div>
    )
    
    
}

export {Employee, AddEmployee, UpdateEmployee, DeleteEmployee};

