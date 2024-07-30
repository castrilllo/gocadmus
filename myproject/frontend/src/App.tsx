import { useState } from 'react';
import logo from './assets/images/logo.png';
import './App.css';
import { ListRunningEC2Instances } from "../wailsjs/go/main/App";

// Define a type for the instance data
type InstanceInfo = {
    id: string;
    name: string;
};

function App() {
    const [instances, setInstances] = useState<InstanceInfo[]>([]);
    const [loading, setLoading] = useState(false);

    const listRunningInstances = () => {
        setLoading(true);
        ListRunningEC2Instances()
            .then((result: InstanceInfo[]) => {
                setInstances(result);
                setLoading(false);
            })
            .catch(err => {
                console.error("Error fetching instances:", err);
                setInstances([]);
                setLoading(false);
            });
    };

    return (
        <div id="App">
            <img src={logo} id="logo" alt="logo" />
            <div id="input" className="input-box">
                <button className="btn" onClick={listRunningInstances} disabled={loading}>
                    {loading ? "Loading..." : "EC2"}
                </button>
            </div>
            <div id="result" className="result">
                {instances.length > 0 ? (
                    <table>
                        <thead>
                            <tr>
                                <th>ID</th>
                                <th>Name</th>
                            </tr>
                        </thead>
                        <tbody>
                            {instances.map(instance => (
                                <tr key={instance.id}>
                                    <td>{instance.id}</td>
                                    <td>{instance.name}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                ) : (
                    <p></p>
                )}
            </div>
        </div>
    );
}

export default App;
