import type { BaseWorkflowProps, WorkflowConfiguration } from "@clutch-sh/core";
export interface WorkflowProps extends BaseWorkflowProps {
}
declare const register: () => WorkflowConfiguration;
export default register;
