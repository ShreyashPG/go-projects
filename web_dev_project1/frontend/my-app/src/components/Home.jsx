import React , {useState} from 'react';


 const Home= ()=>{

    const [greet,setGreet] = useState('')

    const handleClick = async() => {
           try { 
            const response =await fetch("http://localhost:3001")
            const data = await response.text()
            setGreet(data)
           } catch (error) {
               console.log(error)
           }
           
    }
           
    return (
        <div>
            <h1>Home Page</h1>
            <button onClick={handleClick}>Greet</button>
            <h1>{greet}</h1>
        </div>
    )


}

export default Home;