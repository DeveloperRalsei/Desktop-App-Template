import { Button, Group, Space, Stack, Title } from "@mantine/core";
import { createFileRoute } from "@tanstack/react-router";
import { useEffect, useState } from "react";
import { Counter } from "#/app/App";

export const Route = createFileRoute("/")({
    component: RouteComponent,
});

const increase = async () => await Counter("increase");
const decrease = async () => await Counter("decrease");
const getValue = async () => await Counter("");

function RouteComponent() {
    const [counter, setCounter] = useState(0);

    useEffect(() => {
        (async () => {
            setCounter(await getValue());
        })();
    }, [counter]);

    return (
        <Stack h="100vh">
            <Space h="3vh" />

            <Stack h="100%" align="center" justify="center">
                <Title order={3}>Counter: {counter}</Title>
                <Group>
                    <Button
                        color="green"
                        onClick={async () => setCounter(await increase())}
                    >
                        Increase
                    </Button>
                    <Button
                        color="red"
                        onClick={async () => setCounter(await decrease())}
                    >
                        Decrease
                    </Button>
                </Group>
            </Stack>
        </Stack>
    );
}
