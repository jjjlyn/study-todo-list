import * as React from "react"
import {
  ChakraProvider,
  Box,
  VStack,
  theme,
  Spacer,
  StackDivider,
} from "@chakra-ui/react"

import InputForm from "./components/Form";
import {useCallback, useEffect, useState} from "react";
import { Todo } from "./models/todo";
import TodoItem from "./components/Item";
import {changeCompleteTodo, createTodo, deleteTodo, fetchTodo} from "./lib/api/todo";

export default function App() {
  const [items, setItems] = useState<Todo[]>([]);

  const refresh = useCallback(() => {
    fetchTodo().then(value => setItems(value.data.items));
  }, []);

  useEffect(() => {
    refresh();
  }, [refresh]);

  const handleAddItem = useCallback((content: string) => {
    createTodo(content)
        .then(refresh);
  }, [refresh]);

  const handleDeleteItem = useCallback((id: number) => {
    deleteTodo(id)
        .then(refresh);
  }, [refresh]);

  const handleChangeComplete = useCallback((id: number, checked: boolean) => {
    changeCompleteTodo(id, checked)
        .then(refresh);
  }, [refresh]);

  const { location } = window;

  console.log(location.protocol, location.host);

  return (
      <ChakraProvider theme={theme}>
        <VStack h="100%">
          <Spacer />
          <Box
              borderWidth="1px"
              borderRadius="lg"
              p={"2"}
          >
            <InputForm
                addTodo={handleAddItem}
            />
            <VStack
                w="xl"
                h="2xl"
                align="stretch"
                divider={<StackDivider borderColor="gray.200" />}
                overflow="scroll"
            >

              {
                items.map((value) =>
                    <TodoItem
                        key={value.id}
                        todo={value}
                        onDelete={() => handleDeleteItem(value.id)}
                        onChangeComplete={isComplete => handleChangeComplete(value.id, isComplete)}
                    />)
              }
            </VStack>
          </Box>
          <Spacer />
        </VStack>
      </ChakraProvider>
  );
}