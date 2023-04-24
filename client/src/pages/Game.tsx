import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom'
import { socket } from '../socket';

const Game = () => {
  const { id } = useParams()
  const [isConnected, setIsConnected] = useState(socket.connected);

  useEffect(() => {
    console.log("yi")
    function onConnect() {
      console.log("connected")
      setIsConnected(true)
    }

    function onDisconnect() {
      console.log("disconnected")
      setIsConnected(false);
    }

    socket.on('connect', onConnect);
    socket.on('disconnect', onDisconnect);

    socket.open()

    return () => {
      socket.off('connect', onConnect);
      socket.off('disconnect', onDisconnect);
      socket.close()
    };
  }, [])

  return (
    <div>
      <p>Game Code: {id}</p>
      <p>{isConnected}</p>
    </div>
  )
}

export default Game