import apiClient from "./apiClient";
import {Todo} from "../../models/todo";


type FetchResponse = {
    items: Todo[];
}

export function fetchTodo() {
    return apiClient.get<FetchResponse>("/todo");
}

export function createTodo(content: string) {
    return apiClient.post("/todo", {
        content,
    });
}

export function changeCompleteTodo(id: number, isComplete: boolean) {
    return apiClient.patch(`/todo/${id}`, {
        isComplete,
    });
}

export function deleteTodo(id: number) {
    return apiClient.delete(`/todo/${id}`)
}