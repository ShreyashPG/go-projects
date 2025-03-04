import React ,{useState} from 'react';
const HelloWorld = () => {

    const [greet,setGreet] = useState('')
    const handleSubmit = async(e) => {
        e.preventDefault()
        const formData = new FormData(e.target)
        const data = new URLSearchParams(formData)
        try {
            const response = await fetch("http://localhost:3001/hello-world",{
                method: "POST",
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded'
                },
                body: data
            })
            const responseData = await response.text()
            setGreet(responseData)
        } catch (error) {
            console.log(error)
        }
    }

    return (

        <div>
            <form id="greetForm" method="POST" action="http://localhost:3001/hello-world" onSubmit={handleSubmit}>
        <label for="first" >enter greet message</label>
        <input type="text" id="first" name="first"required/>
        <button type="submit">Send</button>
            </form>

            <h1 id="greet">{greet}</h1>
        </div>
    )


}