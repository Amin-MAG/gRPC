package ir.mag.grpc.echo;

public class GrpcThread {

	private Task task;
	private Callback callback;

	public GrpcThread(Task task, Callback callback) {
		this.task = task;
		this.callback = callback;
	}

	public void start() {
		new Thread(() -> {
			GrpcThread.this.task.process();
			GrpcThread.this.callback.post();
		}).start();
	}

	interface Task {
		void process();
	}

	interface Callback {
		void post();
	}
}
