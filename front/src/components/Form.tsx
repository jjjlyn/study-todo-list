import { HStack, Input, IconButton } from "@chakra-ui/react";
import { AddIcon } from '@chakra-ui/icons'
import React, {FormEvent, useCallback, useRef} from "react";

type Props = {
    addTodo: (todo: string) => void;
};

export default function InputForm({ addTodo }: Props) {
    const inputRef = useRef<HTMLInputElement>(null);

    const handleAddTodo = useCallback((todo: string) => {
        if (todo.length === 0) return;
        addTodo(todo);
    }, [addTodo]);

    const handleSubmit = useCallback((event: FormEvent<HTMLFormElement>) => {
        event.preventDefault();
        if (inputRef.current === null) return;
        handleAddTodo(inputRef.current.value);
        inputRef.current.value = "";
    }, [handleAddTodo]);

    return (
        <form onSubmit={handleSubmit}>
            <HStack w="100%">
                <Input
                    ref={inputRef}
                    placeholder="해야 할일"
                />
                <IconButton
                    type="submit"
                    colorScheme="teal"
                    aria-label="추가"
                    icon={<AddIcon />}
                />
            </HStack>
        </form>
    );
}