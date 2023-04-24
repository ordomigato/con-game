import axios, { AxiosRequestConfig } from 'axios';

const baseUrl = "http://localhost:8080/api"

interface CreateGameRequest {
    gameConfig: {
        maxPlayers: number
    }
}

interface JoinGameRequest {
    name: string
    roomCode: string
}

interface GameConfig {
    roomCode: string
    maxPlayers: number
}

interface Player {
    name: string
}

interface Game {
    gameConfig: GameConfig
    players: Player[]
}

const config: AxiosRequestConfig = {
    headers: {
        "Content-Type": "application/json"
      },
}

export async function createGame(payload: CreateGameRequest): Promise<Game> {
    const { data } = await axios.post(`${baseUrl}/create`, payload, config)
    return data
}

export async function joinGame(payload: JoinGameRequest): Promise<Game> {
    const { data } = await axios.post(`${baseUrl}/join`, payload, config)
    return data
}