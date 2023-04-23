import logo from './logo.svg';
import './App.css';
import {useState} from "react";

function App() {
  const [video, setVideo] = useState();
 //
  async function videoFetch(){
    const res = await fetch("http://127.0.0.1:3000/video");
    const data = await res.json();
    console.log(data);
    setVideo(data);
  }
  return (
    <div className="App">
      <video width="1280" height="720" controls="true" src={video} type="video/mp4">
        
      </video>
      <button onClick={videoFetch}>Fetch a video</button>
    </div>
  );
}

export default App;
