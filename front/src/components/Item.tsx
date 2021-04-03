import { Todo } from "../models/todo";
import { Box, HStack, Text, Checkbox, Spacer, IconButton } from "@chakra-ui/react";
import React from "react";
import { DeleteIcon } from "@chakra-ui/icons";

type Props = {
    todo: Todo;
    onDelete: () => void;
    onChangeComplete: (isComplete: boolean) => void;
};

export default function TodoItem({ todo, onDelete, onChangeComplete }: Props) {
    const checked = Boolean(todo.completeAt);
    return (
        <Box
            m={2}
        >
            <HStack w="100%" align={"stretch"}>
                <Checkbox
                    colorScheme="green"
                    isChecked={checked}
                    onChange={event => onChangeComplete(event.target.checked)}
                />
                <Box
                    w="100%"
                    as="button"
                    onClick={() => onChangeComplete(!checked)}
                    textAlign={"start"}
                >
                    <Text
                        as={checked ? "s" : undefined}
                        color={checked ? "gray.300" : "gray.600"}
                    >{todo.content}</Text>
                </Box>
                <Spacer />
                <IconButton
                    colorScheme="gray"
                    aria-label="delete"
                    icon={<DeleteIcon />}
                    onClick={onDelete}
                />
            </HStack>
        </Box>
    );
}