import "./css/index.css";
import "./css/global_fallback.css";
import { Routes, Route } from "react-router-dom";
import Home from "./Routes/Home";
import NavBar from "./Components/NavBar";
import Store from './Routes/Store';

const App = () => {    
    
    return (
        <>
            <NavBar />
            <Routes>                
                
                <Route 
                    path="/" 
                    element={<Home />} 
                />

                <Route 
                    path="/Home" 
                    element={<Home />}
                />

                <Route
                    path="/Store"
                    element={<Store />}
                />
           </Routes>
        </>
    )
}

export default App;