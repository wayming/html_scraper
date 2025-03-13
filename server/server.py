import grpc
from concurrent import futures
import time
import scrape_pb2
import scrape_pb2_grpc
import json

# 实现 gRPC 服务
class ScraperService(scrape_pb2_grpc.HtmlScraperServicer):
    def ProcessPage(self, request, context):
        print(request)
        html_text = request.html_text
        page_type = request.page_type
        
        # 根据页面类型解析 HTML（这里只是一个简单的示例）
        if page_type == "income_statement":
            json_data = {"parsed_data": "Income Statement Data"}
        elif page_type == "balance_sheet":
            json_data = {"parsed_data": "Balance Sheet Data"}
        elif page_type == "cash_flow":
            json_data = {"parsed_data": "Cash Flow Data"}
        else:
            json_data = {"error": "Unknown page type"}

        # 返回解析后的 JSON 数据
        return scrape_pb2.Response(status="OK", json_data=json.dumps(json_data))

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    scrape_pb2_grpc.add_HtmlScraperServicer_to_server(ScraperService(), server)
    server.add_insecure_port('[::]:50051')
    print("Starting server on port 50051...")
    server.start()
    try:
        while True:
            time.sleep(86400)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()
