import HelloWorld from "./hello-world";
const register = () => {
    return {
        developer: {
            name: "daniel",
            contactUrl: "mailto:Daniel@example.com",
        },
        path: "amiibo",
        group: "Amiibo",
        displayName: "Amiibo",
        routes: {
            landing: {
                path: "/",
                displayName: "Amiibo",
                description: "Haha wtf.",
                component: HelloWorld,
            },
        },
    };
};
export default register;
