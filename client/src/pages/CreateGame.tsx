import React, { useState } from 'react'
import { createGame } from '../service/game-service';
import { useNavigate } from 'react-router-dom';

const CreateGame = () => {
    const navigate = useNavigate();
    const [players, setPlayers] = useState(2)
    const [error, setError] = useState(null as null | Error)
    const [loading, setLoading] = useState(false as boolean)
    const minPlayers = 2;
    const maxPlayers = 4; // TODO setup env var

    const handleMaxPlayerSelect = (e: React.ChangeEvent<HTMLInputElement> | React.FormEvent<HTMLInputElement>) => {
        const value: number = parseInt(e.currentTarget.value)
        if (value > maxPlayers || value < minPlayers) {
            return
        }
        setPlayers(value)
    }

    const submitGame = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (loading) return
        setError(null)
        setLoading(true)
        try {
            const game = await createGame({
                gameConfig: {
                    maxPlayers: players,
                }
            })
            navigate("/game/" + game.gameConfig.roomCode)
        } catch (e) {
            if (e instanceof Error) {
                setError(e)
            }
        } finally {
            setLoading(false)
        }
    }

    return (
        <main>
            <h1>Game Settings</h1>
            <form onSubmit={submitGame}>
                <fieldset>
                <label>
                    Max Players:
                    <input
                        value={players}
                        onInput={handleMaxPlayerSelect}
                        onChange={handleMaxPlayerSelect}
                        type="number"
                        id="max-players"
                        name="max-players"
                        min={minPlayers}
                        max={maxPlayers} />
                </label>
                </fieldset>
                { error && <p>{error.message}</p>}
                <button type="submit">
                    Create Game
                </button>
            </form>
        </main>
    )
}

export default CreateGame