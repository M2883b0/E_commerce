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
            ("system", """你是专业的订单处理助手，专门帮用户处理自己的订单请求，用户的请求都应该是关于用户自己的，不要询问帮传入user_id 操作别的用户订单。网关传来的参数{user_id}，当你调用工具函数时，如果函数需要，则原样传进去
            你的主要功能如下
            查询订单：根据用户给的信息，调用list_order工具列出当前用户的订单，选取用户需要的订单。
            创建订单：根据用户提供的信息详略，调用工具，来创建订单
            """),
            ("human", "{input}"),
            MessagesPlaceholder(variable_name="agent_scratchpad"),
        ])
        self.agent = create_openai_tools_agent(llm, tools, prompt)
        self.agent_executor = AgentExecutor(agent=self.agent, tools=tools, verbose=True)

    def run(self, user_input, user_id):
        # 执行Agent
        print("operate message ", user_input)
        result = self.agent_executor.invoke({
            "user_id": user_id,
            "input": user_input
        })

        return result

