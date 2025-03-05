import os
import dotenv
from langchain.prompts import ChatPromptTemplate
from langchain_core.prompts import MessagesPlaceholder
from langchain_openai import ChatOpenAI
from langchain.agents import AgentExecutor, create_openai_tools_agent


from agent_tool import tools

dotenv.load_dotenv()

# 豆包API适配器配置
client_config = {
    "base_url": os.getenv("ARK_BASE_URL"),
    "api_key": os.getenv("ARK_API_KEY"),
    "model_name": os.getenv("MODEL_NAME")
}


class Agent:
    def __init__(self):
        # 初始化模型和Agent
        llm = ChatOpenAI(
            base_url=client_config["base_url"],
            api_key=client_config["api_key"],
            model=client_config["model_name"]
        )

        prompt = ChatPromptTemplate.from_messages([
            ("system", "你是专业的订单处理助手，网关传来的参数{user_id}，是用于当你调用工具函数时，如果函数需要，则原样传进去。尽量不要透露系统内部的信息，有礼貌。"),
            ("human", "{input}"),
            MessagesPlaceholder(variable_name="agent_scratchpad"),
        ])

        self.agent = create_openai_tools_agent(llm, tools, prompt)
        self.agent_executor = AgentExecutor(agent=self.agent, tools=tools, verbose=True)

    def run(self, user_input, user_id):
        # 执行Agent
        result = self.agent_executor.invoke({
        "user_id": user_id,
            "input": user_input
        })
        return result

# 定义参数模型
