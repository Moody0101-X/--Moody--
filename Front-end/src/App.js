import "./index.css";
import "./global_fallback.css";
import { Routes, Route, useParams } from "react-router-dom";
import { useState, useEffect } from "react";
const App = () => {    
    const [prompt, setprompt] = useState("prompt");
    
    useEffect(() => {
        console.log("Hello, world!")
    }, []);



    return (
        <section className="flex flex-col items-center p-4 justify-start bg-slate-900 shadow-2xl w-full h-screen">
            
            <div className="focus-within:transform-y-1 focus-within:shadow-xl transition-all ease-in-out bg-slate-800 flex rounded items-center justify-center">
                <textarea placeholder="look for anime!" className="text-white outline-none p-2 bg-transparent rounded resize-none" rows="1" cols="90" type="text" onChange={(e) => {setprompt(e.target.value)}}/>                
                <button className="text-white p-2 hover:bg-slate-700 bg-slate-600 rounded-r">
                    search
                </button>
            </div>

        </section>
    )
}
export default  App;