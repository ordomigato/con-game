import React from 'react'
import { Link } from 'react-router-dom'

const Home = () => {
  return (
    <main>
        <h1>Welcome!</h1>
        <div>
            <div>
                <Link className="btn" to="/create">Host Game</Link>
            </div>
            <div>
                <Link className="btn" to="/join">Join Game</Link>
            </div>
        </div>
    </main>
  )
}

export default Home